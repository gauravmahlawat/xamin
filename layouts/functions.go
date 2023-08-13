package main

import (
    "github.com/gohugoio/hugo/identity"
    "github.com/gohugoio/hugo/tpl/internal"
)

func init() {
    identity.AddFuncMap("randomBlog", func(limit int) interface{} {
        return func() internal.TemplateFuncResult {
            site := internal.GetSiteFromContext()

            if site == nil {
                return internal.TemplateFuncError("No site in .Site.RegularPages")
            }

            randomIndex := site.Random.Intn(site.RegularPages.Len())

            return site.RegularPages[randomIndex]
        }
    })
}
