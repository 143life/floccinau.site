const matrixEl = document.getElementById('matrix')

function BuildMatrix(cols, rows) {
	const frag = document.createDocumentFragment();
	for (let r = 0; r < rows; r++) {
		for (let c = 0; c < cols; c++) {
			const cell = document.createElement('button');
			cell.type = 'button';
			cell.className = 'cell';
			cell.setAttribute('role', 'gridcell');
			cell.setAttribute('aria-label', `Ячейка ${r + 1}-${c + 1}`);
			cell.addEventListener('click', () => {
				cell.style.backgroundColor = randomColor();
				cell.style.boxShadow = '0 6px 18px rgba(0, 0, 0,.10)';
			});
			cell.addEventListener('contextmenu', (e) => {
				e.preventDefault();
				cell.style.backgroundColor = '';
				cell.style.boxShadow = '';
			});
			frag.appendChild(cell);
		}
	}
	matrixEl.innerHTML = '';
	matrixEl.appendChild(frag);
}

function randomColor() {
	const h = Math.floor(Math.random() * 360);
	const s = 70 + Math.floor(Math.random() * 30);
	const l = 55 + Math.floor(Math.random() * 15);
	return `hsl(${h} ${s} ${l}%)`;
}

const cols = parseInt(matrixEl.dataset.cols || '20', 10);
const rows = parseInt(matrixEl.dataset.rows || '10', 10);
BuildMatrix(cols, rows);

// Optional: keyboard toggle on focused cell (space/enter)
matrixEl.addEventListener('keydown', (e) => {
  if (e.key === ' ' || e.key === 'Enter') {
    const el = document.activeElement;
    if (el && el.classList.contains('cell')) {
      e.preventDefault();
      el.style.backgroundColor = randomColor();
      el.style.boxShadow = '0 6px 18px rgba(0,0,0,.10)';
    }
  }
});