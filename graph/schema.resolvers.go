package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"
	"gqlgen-playground/graph/gql"
)

// Author is the resolver for the author field.
func (r *issueResolver) Author(ctx context.Context, obj *gql.Issue) (*gql.User, error) {
	// 1. Loaderに検索条件となるIDを登録(この時点では即時実行されない)
	thunk := r.Loaders.UserLoader.Load(ctx, obj.Author.ID)
	// 2. LoaderがDBに対してデータ取得処理を実行するまで待って、結果を受け取る
	user, err := thunk()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// AddProjectV2ItemByID is the resolver for the addProjectV2ItemById field.
func (r *mutationResolver) AddProjectV2ItemByID(ctx context.Context, input gql.AddProjectV2ItemByIDInput) (*gql.AddProjectV2ItemByIDPayload, error) {
	panic(fmt.Errorf("not implemented: AddProjectV2ItemByID - addProjectV2ItemById"))
}

// Owner is the resolver for the owner field.
func (r *projectV2Resolver) Owner(ctx context.Context, obj *gql.ProjectV2) (*gql.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// Repository is the resolver for the repository field.
func (r *queryResolver) Repository(ctx context.Context, name string, owner string) (*gql.Repository, error) {
	panic(fmt.Errorf("not implemented: Repository - repository"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*gql.User, error) {
	return r.Service.GetUserByName(ctx, name)
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (gql.Node, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Owner is the resolver for the owner field.
func (r *repositoryResolver) Owner(ctx context.Context, obj *gql.Repository) (*gql.User, error) {
	panic(fmt.Errorf("not implemented: Owner - owner"))
}

// Issue is the resolver for the issue field.
func (r *repositoryResolver) Issue(ctx context.Context, obj *gql.Repository, number int) (*gql.Issue, error) {
	panic(fmt.Errorf("not implemented: Issue - issue"))
}

// Issues is the resolver for the issues field.
func (r *repositoryResolver) Issues(ctx context.Context, obj *gql.Repository, after *string, before *string, first *int, last *int) (*gql.IssueConnection, error) {
	panic(fmt.Errorf("not implemented: Issues - issues"))
}

// PullRequest is the resolver for the pullRequest field.
func (r *repositoryResolver) PullRequest(ctx context.Context, obj *gql.Repository, number int) (*gql.PullRequest, error) {
	panic(fmt.Errorf("not implemented: PullRequest - pullRequest"))
}

// PullRequests is the resolver for the pullRequests field.
func (r *repositoryResolver) PullRequests(ctx context.Context, obj *gql.Repository, after *string, before *string, first *int, last *int) (*gql.PullRequestConnection, error) {
	panic(fmt.Errorf("not implemented: PullRequests - pullRequests"))
}

// Issue returns gql.IssueResolver implementation.
func (r *Resolver) Issue() gql.IssueResolver { return &issueResolver{r} }

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// ProjectV2 returns gql.ProjectV2Resolver implementation.
func (r *Resolver) ProjectV2() gql.ProjectV2Resolver { return &projectV2Resolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

// Repository returns gql.RepositoryResolver implementation.
func (r *Resolver) Repository() gql.RepositoryResolver { return &repositoryResolver{r} }

type issueResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type projectV2Resolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type repositoryResolver struct{ *Resolver }
