function loadMovies() {
    fetch('http://localhost:8081/movies')
        .then(res => res.json())
        .then(data => {
            const moviesDiv = document.getElementById('movies');
            moviesDiv.innerHTML = '';

            if (data.length === 0) {
                moviesDiv.innerHTML = '<p>Фильмы пока не добавлены.</p>';
                return;
            }

            data.forEach(movie => {
                const card = document.createElement('div');
                card.className = 'movie-card';
                card.innerHTML = `
                    <h3>${movie.title}</h3>
                    <p><strong>Жанр:</strong> ${movie.genre}</p>
                    <p><strong>Длительность:</strong> ${movie.duration} мин</p>
                    <p>${movie.description}</p>
                    <button class="delete-btn" data-id="${movie.id}">Удалить</button>
                `;

                moviesDiv.appendChild(card);
            });

            // 🔥 Привязываем события на все кнопки после генерации
            document.querySelectorAll(".delete-btn").forEach(btn => {
                btn.addEventListener("click", () => {
                    const id = btn.dataset.id;
                    deleteMovie(id);
                });
            });
        })
        .catch(err => {
            console.error('Ошибка при получении фильмов:', err);
        });
}

// 🎬 Обработчик формы добавления
document.getElementById('movie-form').addEventListener('submit', function (e) {
    e.preventDefault();

    const movie = {
        title: document.getElementById('title').value,
        description: document.getElementById('description').value,
        genre: document.getElementById('genre').value,
        duration: parseInt(document.getElementById('duration').value)
    };

    fetch('http://localhost:8081/movies', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(movie)
    })
        .then(res => res.json())
        .then(() => {
            loadMovies();
            this.reset();
        })
        .catch(err => console.error('Ошибка при добавлении фильма:', err));
});

// 🧨 Удаление фильма
function deleteMovie(id) {
    if (!confirm("Удалить фильм?")) return;

    fetch(`http://localhost:8081/movies/${id}`, {
        method: 'DELETE'
    })
        .then(res => {
            if (!res.ok) throw new Error('Ошибка при удалении');
            return res.json();
        })
        .then(() => loadMovies())
        .catch(err => console.error('Ошибка при удалении фильма:', err));
}

// 🚀 Загружаем фильмы при старте страницы
loadMovies();
