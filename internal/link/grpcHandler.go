package link

import (
	"context"

	"github.com/user-xat/short-link/internal/models"
	"github.com/user-xat/short-link/pkg/req"
	pb "github.com/user-xat/short-link/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"gorm.io/gorm"
)

type GRPCHandlerDeps struct {
	LinkRepository ILinkRepository
}

type GRPCHandler struct {
	pb.UnimplementedShortLinkServer
	LinkRepository ILinkRepository
}

func NewGRPCHandler(deps GRPCHandlerDeps) *GRPCHandler {
	return &GRPCHandler{
		LinkRepository: deps.LinkRepository,
	}
}

func (g *GRPCHandler) Create(ctx context.Context, url *wrapperspb.StringValue) (*pb.Link, error) {
	reqLink := LinkCreateRequest{Url: url.Value}
	err := req.IsValid(reqLink)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	link := models.NewLink(reqLink.Url)
	for range 100 {
		var existedLink *models.Link
		existedLink, err = g.LinkRepository.GetByHash(link.Hash)
		if existedLink == nil {
			break
		}
		link.GenerateHash()
	}
	if err == nil {
		return nil, status.Error(codes.Internal, "Can't create unique hash for link. Try again or contact us.")
	}
	createdLink, err := g.LinkRepository.Create(link)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error to create %s link: %v", link.Url, err)
	}

	return &pb.Link{
		Id:   uint64(createdLink.ID),
		Url:  createdLink.Url,
		Hash: createdLink.Hash,
	}, status.Error(codes.OK, "")
}

func (g *GRPCHandler) GetByHash(ctx context.Context, hash *wrapperspb.StringValue) (*pb.Link, error) {
	link, err := g.LinkRepository.GetByHash(hash.Value)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error to get record by hash %s: %v", hash.Value, err)
	}

	return &pb.Link{
		Id:   uint64(link.ID),
		Url:  link.Url,
		Hash: link.Hash,
	}, status.Error(codes.OK, "")
}

func (g *GRPCHandler) GetById(ctx context.Context, id *wrapperspb.UInt64Value) (*pb.Link, error) {
	link, err := g.LinkRepository.GetById(uint(id.Value))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error to get record by id %d: %v", id.Value, err)
	}
	return &pb.Link{
		Id:   uint64(link.ID),
		Url:  link.Url,
		Hash: link.Hash,
	}, status.Error(codes.OK, "")
}

func (g *GRPCHandler) GetAll(ctx context.Context, option *pb.LimitOffset) (*pb.Links, error) {
	links := g.LinkRepository.GetAll(int(option.Limit), int(option.Offset))
	res := make([]*pb.Link, len(links))
	for i := range links {
		res[i] = &pb.Link{
			Id:   uint64(links[i].ID),
			Url:  links[i].Url,
			Hash: links[i].Hash,
		}
	}
	return &pb.Links{
		Link: res,
	}, status.Error(codes.OK, "")
}

func (g *GRPCHandler) Update(ctx context.Context, link *pb.Link) (*pb.Link, error) {
	updatedLink, err := g.LinkRepository.Update(&models.Link{
		Model: gorm.Model{
			ID: uint(link.Id),
		},
		Url:  link.Url,
		Hash: link.Hash,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Link{
		Id:   uint64(updatedLink.ID),
		Url:  updatedLink.Url,
		Hash: updatedLink.Hash,
	}, status.Error(codes.OK, "")
}

func (g *GRPCHandler) Delete(ctx context.Context, id *wrapperspb.UInt64Value) (*pb.Void, error) {
	err := g.LinkRepository.Delete(uint(id.Value))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.Void{}, status.Error(codes.OK, "")
}

func (g *GRPCHandler) Count(ctx context.Context, void *pb.Void) (*wrapperspb.UInt64Value, error) {
	cnt := g.LinkRepository.Count()
	return &wrapperspb.UInt64Value{
		Value: uint64(cnt),
	}, status.Error(codes.OK, "")
}
