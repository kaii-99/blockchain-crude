package keeper

import (
	"context"
	"errors"

	"crude/x/crude/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	// Ensure the tags field is not empty
	if len(msg.Tags) == 0 {
		return nil, errors.New("tags are required")
	}
	
	var post = types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
		Tags:    msg.Tags,  // Save the tags to the post
	}
	id := k.AppendPost(
		ctx,
		post,
	)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
