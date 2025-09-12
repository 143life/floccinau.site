const windowInnerWidth = window.innerWidth;
const windowInnerHeight = window.innerHeight;
const count = windowInnerWidth / 30;

const container = document.querySelector('.pixel-matrix');
const pixelHTML = `<div class="pixel"></div>`;
const pixelsHTML = pixelHTML.repeat(count);
const pixelRowHTML = `<div class="pixel-row">` + pixelsHTML + `</div>`
const pixelMatrixHTML = pixelRowHTML.repeat(count / 4)
container.innerHTML = pixelMatrixHTML;

