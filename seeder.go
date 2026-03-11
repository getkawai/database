package database

import (
	"context"
	_ "embed"
	"fmt"
	"log"
)

//go:embed seed_data.sql
var seedDataSQL string

// SeedAvailableAgents loads pre-seeded agents from SQL dump
func (s *Service) SeedAvailableAgents(ctx context.Context) error {
	// 1. Check if we already have agents
	var count int
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM agents WHERE id != ?", DefaultInboxAgentID).Scan(&count)
	if err != nil {
		log.Printf("Failed to check existing agents: %v", err)
	}

	// Skip if we already have agents (more than just inbox)
	if count > 0 {
		log.Printf("Agents already seeded (%d agents found). Skipping.", count)
		return nil
	}

	log.Println("Loading pre-seeded agents from SQL dump...")

	// 2. Execute embedded SQL dump
	_, err = s.db.ExecContext(ctx, seedDataSQL)
	if err != nil {
		return fmt.Errorf("failed to execute seed data: %w", err)
	}

	// 3. Count inserted agents
	err = s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM agents WHERE id != ?", DefaultInboxAgentID).Scan(&count)
	if err == nil {
		log.Printf("✅ Successfully loaded %d agents from seed data", count)
	}

	return nil
}
