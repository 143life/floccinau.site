Floccinau Site
===============

A small personal website written in Go with server‑rendered HTML templates and a tiny bit of vanilla JS/CSS for a pixel matrix demo.

Features
--------
- Simple HTTP server in Go
- HTML templates with partials (`nav`, `pixel-art`)
- Static assets served under `/static`
- Accessible, interactive pixel matrix (keyboard and mouse)

Tech Stack
----------
- Go (net/http, html/template)
- HTML/CSS/JS (no frameworks)

Project Structure
-----------------
```text
cmd/web/                # App entrypoint and routing
  main.go               # Server setup and static file handler
  handlers.go           # HTTP handlers, template rendering
ui/
  html/
    base.tmpl.html      # Base layout
    pages/
      home.tmpl.html    # Home page
    partials/
      nav.tmpl.html     # Navigation
      pixel-art.tmpl.html  # Pixel matrix section and script include
  static/
    css/
      main.css          # Global styles
      pixel-art.css     # Pixel matrix styles
      stylesheet.css    # Font-face, etc.
    js/
      pixel-art.js      # Matrix generation logic
    fonts/              # Web fonts (woff2)
```

Getting Started
---------------
Prerequisites:
- Go 1.21+ installed

Run the development server:
```bash
go run ./cmd/web
```
Open the site at `http://localhost:4000`.

Static Assets
-------------
- Served from `./ui/static` at `/static`.
- Content types are set by a small middleware in `cmd/web/main.go`.

Templates
---------
- Base layout: `ui/html/base.tmpl.html` includes global CSS and partials.
- Home page content: `ui/html/pages/home.tmpl.html`.
- Pixel matrix partial adds the container and loads `pixel-art.js`.

Pixel Matrix Customization
--------------------------
- Grid container is `#matrix` (role="grid"). Cells are `button.cell` (role="gridcell").
- Change grid size via data attributes in `pixel-art.tmpl.html`:
```html
<div id="matrix" class="matrix" role="grid" data-cols="20" data-rows="10"></div>
```
- Keyboard support: focus a cell and press Space/Enter to toggle color.


Troubleshooting
---------------
- Only one long row appears:
  - Ensure `matrixEl.innerHTML = ''` and `appendChild(frag)` are executed once after both loops in `pixel-art.js`.
  - Verify `.matrix { display: grid; }` is applied and CSS is loaded.
- Static files not loading:
  - Confirm links like `/static/css/main.css` resolve. The server mounts `./ui/static` at `/static`.
- Font not displaying:
  - Make sure `PressStart2P-Regular.woff2` exists and the path in `stylesheet.css` is correct.

Configuration
-------------
- Server port is set in `cmd/web/main.go` (`:4000`). Change if needed.
- Add new pages by creating a file in `ui/html/pages/` and rendering it from a new handler.

License
-------
MIT
# Start of my site development

## README not ready, sorry :)