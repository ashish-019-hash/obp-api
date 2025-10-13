package services

import (
	"time"

	"github.com/ashish-019-hash/obp-api-backend/internal/models"
	"gorm.io/gorm"
)

type SessionService struct {
	db            *gorm.DB
	configService *ConfigService
	randomService *SecureRandomService
}

func NewSessionService(db *gorm.DB, configService *ConfigService, randomService *SecureRandomService) *SessionService {
	return &SessionService{
		db:            db,
		configService: configService,
		randomService: randomService,
	}
}

func (ss *SessionService) CreateAPISession(userID, consumerID, ipAddress, userAgent string) (*models.APISession, error) {
	timeoutMinutes := ss.configService.GetConfigInt("api.session.timeout.minutes", 60)
	
	session := models.NewAPISession(userID, consumerID, ipAddress, userAgent, timeoutMinutes)
	
	if err := ss.db.Create(session).Error; err != nil {
		return nil, err
	}
	
	return session, nil
}

func (ss *SessionService) GetAPISession(sessionID string) (*models.APISession, error) {
	var session models.APISession
	err := ss.db.Where("session_id = ? AND is_active = ? AND expires_at > ?", 
		sessionID, true, time.Now()).First(&session).Error
	
	if err != nil {
		return nil, err
	}
	
	session.LastAccessAt = time.Now()
	ss.db.Save(&session)
	
	return &session, nil
}

func (ss *SessionService) RefreshAPISession(sessionID string) error {
	timeoutMinutes := ss.configService.GetConfigInt("api.session.timeout.minutes", 60)
	
	return ss.db.Model(&models.APISession{}).
		Where("session_id = ? AND is_active = ?", sessionID, true).
		Update("expires_at", time.Now().Add(time.Duration(timeoutMinutes)*time.Minute)).Error
}

func (ss *SessionService) InvalidateAPISession(sessionID string) error {
	return ss.db.Model(&models.APISession{}).
		Where("session_id = ?", sessionID).
		Update("is_active", false).Error
}

func (ss *SessionService) InvalidateUserSessions(userID string) error {
	return ss.db.Model(&models.APISession{}).
		Where("user_id = ?", userID).
		Update("is_active", false).Error
}

func (ss *SessionService) CleanupExpiredSessions() error {
	return ss.db.Where("expires_at < ? OR is_active = ?", time.Now(), false).
		Delete(&models.APISession{}).Error
}

func (ss *SessionService) GetUserSessions(userID string) ([]models.APISession, error) {
	var sessions []models.APISession
	err := ss.db.Where("user_id = ? AND is_active = ? AND expires_at > ?", 
		userID, true, time.Now()).Find(&sessions).Error
	
	return sessions, err
}

func (ss *SessionService) CreateUserRefresh(userID string) (*models.UserRefresh, error) {
	expirationHours := ss.configService.GetConfigInt("refresh.token.expiration.hours", 168) // 7 days
	
	refresh := models.NewUserRefresh(userID, expirationHours)
	
	if err := ss.db.Create(refresh).Error; err != nil {
		return nil, err
	}
	
	return refresh, nil
}

func (ss *SessionService) ValidateRefreshToken(refreshToken string) (*models.UserRefresh, error) {
	var refresh models.UserRefresh
	err := ss.db.Where("refresh_token = ? AND is_active = ? AND expires_at > ?", 
		refreshToken, true, time.Now()).First(&refresh).Error
	
	return &refresh, err
}

func (ss *SessionService) InvalidateRefreshToken(refreshToken string) error {
	return ss.db.Model(&models.UserRefresh{}).
		Where("refresh_token = ?", refreshToken).
		Update("is_active", false).Error
}

func (ss *SessionService) InvalidateUserRefreshTokens(userID string) error {
	return ss.db.Model(&models.UserRefresh{}).
		Where("user_id = ?", userID).
		Update("is_active", false).Error
}

func (ss *SessionService) StartSessionCleanupRoutine() {
	cleanupInterval := ss.configService.GetConfigInt("api.session.cleanup.interval.minutes", 15)
	
	ticker := time.NewTicker(time.Duration(cleanupInterval) * time.Minute)
	go func() {
		for range ticker.C {
			if err := ss.CleanupExpiredSessions(); err != nil {
				continue
			}
		}
	}()
}

func (ss *SessionService) GetSessionStats() (map[string]interface{}, error) {
	var totalSessions int64
	var activeSessions int64
	var expiredSessions int64
	
	if err := ss.db.Model(&models.APISession{}).Count(&totalSessions).Error; err != nil {
		return nil, err
	}
	
	if err := ss.db.Model(&models.APISession{}).
		Where("is_active = ? AND expires_at > ?", true, time.Now()).
		Count(&activeSessions).Error; err != nil {
		return nil, err
	}
	
	if err := ss.db.Model(&models.APISession{}).
		Where("expires_at < ?", time.Now()).
		Count(&expiredSessions).Error; err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"total_sessions":   totalSessions,
		"active_sessions":  activeSessions,
		"expired_sessions": expiredSessions,
	}, nil
}
