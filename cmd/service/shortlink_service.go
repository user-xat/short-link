package main

import (
	"context"

	pb "github.com/user-xat/short-link-server/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ShortLinkService struct {
	pb.UnimplementedShortLinkServer
	sl *ShortLink
}

func NewShortLinkService(shortlink *ShortLink) *ShortLinkService {
	return &ShortLinkService{
		sl: shortlink,
	}
}

func (s *ShortLinkService) Add(ctx context.Context, link *wrapperspb.StringValue) (*pb.Link, error) {
	data, err := s.sl.Set(ctx, link.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error to add %s link: %v", link.Value, err)
	}

	return &pb.Link{
		Short:  data.Short,
		Source: data.Source,
	}, status.Error(codes.OK, "")
}

func (s *ShortLinkService) Get(ctx context.Context, short *wrapperspb.StringValue) (*pb.Link, error) {
	data, err := s.sl.GetLink(ctx, short.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error to get %s: %v", short.Value, err)
	}

	return &pb.Link{
		Short:  data.Short,
		Source: data.Source,
	}, status.Error(codes.OK, "")
}
