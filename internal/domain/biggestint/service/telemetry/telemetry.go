package telemetry

import (
	"context"
	"github.com/rs/zerolog/log"
)

type Telemetry struct {
}

func (t *Telemetry) AnnounceInts(ctx context.Context, ints []int) {
	log.Ctx(ctx).Info().Msgf("ints inputted to find the biggest int: %v", ints)
}

func (t *Telemetry) AnnounceBiggestIntFound(ctx context.Context, biggest int) {
	log.Ctx(ctx).Info().Msgf("biggest int found: %v", biggest)
}
