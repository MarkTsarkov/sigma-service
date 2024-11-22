package grpc

import (
	"context"
	"github.com/marktsarkov/sigma-service/internal/entity"
	"github.com/marktsarkov/sigma-service/internal/service"
	pb "github.com/marktsarkov/sigma-service/pkg/note"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NoteServer struct {
	pb.UnimplementedNoteServer
	serv service.NoteService
}

func NewNoteServer(unimplementedNoteServer pb.UnimplementedNoteServer, serv service.NoteService) *NoteServer {
	return &NoteServer{
		UnimplementedNoteServer: unimplementedNoteServer,
		serv:                    serv,
	}
}

func (s *NoteServer) GetById(ctx context.Context, req *pb.GetByIDRequest) (resp *pb.GetByIDResponse, err error) {

	id := req.GetId()
	note, err := s.serv.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.GetByIDResponse{
		Note: &pb.NoteInfo{
			Id: note.ID,
			Content: &pb.NoteContent{
				Title: note.Title,
				Body:  note.Body,
			},
			CreatedAt: timestamppb.New(note.CreatedAt),
			UpdatedAt: timestamppb.New(note.UpdatedAt),
		},
	}, nil
}

func (s *NoteServer) Create(ctx context.Context, req *pb.CreateRequest) (resp *pb.CreateResponse, err error) {
	note := &entity.Note{
		Title: req.Note.Content.Title,
		Body:  req.Note.Content.Body,
	}

	id, err := s.serv.Create(ctx, note)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Id: id,
	}, nil
}
