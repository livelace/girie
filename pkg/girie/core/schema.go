package core

import (
	"github.com/graphql-go/graphql"
)

var articleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Article",
	Fields: graphql.Fields{
		"html": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Context.Value("data").(Data).Article.HTML, nil
			},
		},
		"images": &graphql.Field{
			Type: graphql.NewList(imageType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ExtractImages(p.Context.Value("data").(Data).Article.HTML), nil
			},
		},
		"text": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Context.Value("data").(Data).Article.Text, nil
			},
		},
	},
})

var imageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Image",
	Fields: graphql.Fields{
		"alt": &graphql.Field{
			Type: graphql.String,
		},
		"src": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var pageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Page",
	Fields: graphql.Fields{
		"html": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Context.Value("data").(Data).Page.HTML, nil
			},
		},
		"images": &graphql.Field{
			Type: graphql.NewList(imageType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ExtractImages(p.Context.Value("data").(Data).Page.HTML), nil
			},
		},
		"text": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return SanitizeHTMLTags(p.Context.Value("data").(Data).Page.HTML), nil
			},
		},
	},
})

var dataType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Data",
	Fields: graphql.Fields{
		"html": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Context.Value("data").(Data).HTML, nil
			},
		},
		"url": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Context.Value("data").(Data).URL, nil
			},
		},

		"jsonld": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ExtractJSONLd(p.Context.Value("data").(Data).Page.HTML), nil
			},
		},
		"microdata": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ExtractMicrodata(p.Context.Value("data").(Data).Page.HTML,
					p.Context.Value("data").(Data).URL), nil
			},
		},
		"opengraph": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ExtractOpengraph(p.Context.Value("data").(Data).Page.HTML), nil
			},
		},
		"rdfa": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ExtractRDFA(p.Context.Value("data").(Data).Page.HTML), nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ExtractTitle(p.Context.Value("data").(Data).Page.HTML), nil
			},
		},

		"article": &graphql.Field{
			Type: articleType,
		},
		"page": &graphql.Field{
			Type: pageType,
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"data": &graphql.Field{
			Type:        dataType,
			Description: "Get single todo",
			Args: graphql.FieldConfigArgument{
				"html": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"url": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return Data{}, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})
