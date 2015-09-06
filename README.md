# About

This is a webserver written in go with [IndieWeb](http://indiewebcamp.com) principles in mind.

Currently this does not do much other than serve some static templated files.

## Running

After downloading and compiling the app. You can run the binary by passing it a `webdir` parameter on the commandline. eg. `indigo -webdir=/srv/www` in the webdir directory the following subdirectories are expected; 'templates', 'static', 'content'. Currently templates are hardcoded as page.html and header.html (this portion will be refactored at some point).

### templates directory

As noted above the templates directory only supports page.html and header.html.

#### page.html

````
{{define "page"}}
<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>{{template "title"}}</title>
  <link rel="stylesheet" href="/static/css/foundation.css">
  <link rel="stylesheet" href="/static/foundation-icons/foundation-icons.css">

  <!-- This is how you would link your custom stylesheet -->
  <link rel="stylesheet" href="/static/css/app.css">

  <script src="/static/js/vendor/modernizr.js"></script>

</head>
<body>
  <script src="/static/js/vendor/jquery.js"></script>
  <script src="/static/js/foundation/foundation.js"></script>
  {{template "header"}}
  {{template "body"}}

  <script>$(document).foundation();</script>
</body>
</html>
{{end}}
````

#### header.html
````
{{define "header"}}
<div class="icon-bar one-up">
  <a href="http://jgraham909.com/" class="item"><i class="fi-home"></i><label>Home</label></a>
</div>
{{end}}
````

### static directory

Any file located under static will automatically be served as a static file when accessed as http://example.com/static/PATH-TO-STATIC-FROM-STATIC-DIR

### content directory

Intended to be served as golang templates. Note that any template values used in the templates directory can be used. As an example if using the example page.html and header.html templates from above then you can make a post by naming it how you want to access it. For example if I wanted a blog post at http://example.com/blog/article-1 I would make a directory `blog` inside my content directory and then a corresponding file called `artile-1` (note no file extension). As an example here is a possible template for article-1;

````
{{define "title"}}example.com: Article 1{{end}}

{{define "body"}}
This is my first blog post! Woohoo!
{{end}}

````

## Current Roadmap

1. Add capability to add notes, articles, and replies
2. Add webmention functionality
  * Include a cache of incoming parsed webmention source URLs.
  * On repeat webmentions capture a diff and only notify of changed content.
3. Add Search
4. POSSE to Twitter
5. Documentation

## This code is running http://jgraham909.com
