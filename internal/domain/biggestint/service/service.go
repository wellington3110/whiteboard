package service

import (
	"context"
	"github.com/wellington3110/whiteboard/internal/domain/biggestint"
	"github.com/wellington3110/whiteboard/internal/domain/biggestint/service/telemetry"
	"math"
)

type Announcer interface {
	AnnounceInts(ctx context.Context, ints []int)
	AnnounceBiggestIntFound(ctx context.Context, biggest int)
}

type Service struct {
	announcer Announcer
}

func New() *Service {
	return &Service{
		announcer: &telemetry.Telemetry{},
	}
}

func (s *Service) FindBiggestInt(ctx context.Context, ints []int) (int, error) {
	if len(ints) == 0 {
		return 0, biggestint.NewIntsMustNotBeEmptyError()
	}
	return s.findBiggestInt(ctx, ints)
}

func (s *Service) findBiggestInt(ctx context.Context, ints []int) (int, error) {
	s.announcer.AnnounceInts(ctx, ints)
	if len(ints) == 1 {
		return ints[0], nil
	}
	biggest := math.MinInt64
	for _, n := range ints {
		if n > biggest {
			biggest = n
		}
	}
	s.announcer.AnnounceBiggestIntFound(ctx, biggest)
	return biggest, nil
}
