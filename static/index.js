async function fetchVideos() {
    try {
        const response = await fetch('/api/videos');
        const videos = await response.json();
        displayVideos(videos);
    } catch (error) {
        console.error('Error fetching videos:', error);
    }
}

function displayVideos(videos) {
    const videoList = document.getElementById('videoList');
    videoList.innerHTML = videos.map(video => `
        <div class="video-card">
            <div class="video-preview cursor-pointer" onclick="openModal('/videos/${video}')">
                <video preload="metadata">
                    <source src="/videos/${video}" type="video/mp4">
                </video>
                <div class="absolute inset-0 flex items-center justify-center bg-black bg-opacity-40 opacity-0 hover:opacity-100 transition-opacity">
                    <i class="fas fa-play text-4xl text-white"></i>
                </div>
            </div>
            <div class="p-4 flex-grow flex flex-col justify-between">
                <h3 class="text-lg font-semibold text-gray-200 mb-3">${video.replace('.mp4', '')}</h3>
                <div class="action-buttons">
                    <div class="tooltip">
                        <button onclick="copyVideoUrl('/videos/${video}')" 
                                class="action-button">
                            <i class="fas fa-link"></i>
                        </button>
                        <span class="tooltiptext">URL Kopyala</span>
                    </div>
                    <div class="tooltip">
                        <a href="/videos/${video}" download 
                           class="action-button">
                            <i class="fas fa-download"></i>
                        </a>
                        <span class="tooltiptext">İndir</span>
                    </div>
                    <div class="tooltip">
                        <button onclick="deleteVideo('${video}')" 
                                class="action-button text-red-500 hover:text-red-400">
                            <i class="fas fa-trash"></i>
                        </button>
                        <span class="tooltiptext">Sil</span>
                    </div>
                </div>
            </div>
        </div>
    `).join('');
}

function openModal(videoUrl) {
    const modal = document.getElementById('videoModal');
    const modalVideo = document.getElementById('modalVideo');
    modalVideo.src = videoUrl;
    modal.style.display = 'block';
    setTimeout(() => modal.classList.add('show'), 10);
    modalVideo.play();

    // Video yüklendiğinde en-boy oranını kontrol et
    modalVideo.onloadedmetadata = function() {
        const videoRatio = modalVideo.videoWidth / modalVideo.videoHeight;
        const modalContent = document.querySelector('.modal-content');
        
        if (videoRatio < 1) { // Dikey video
            modalContent.style.width = 'auto';
            modalContent.style.maxWidth = '600px';
        } else { // Yatay video
            modalContent.style.width = '90%';
            modalContent.style.maxWidth = '1200px';
        }
    };
}

function closeModal() {
    const modal = document.getElementById('videoModal');
    const modalVideo = document.getElementById('modalVideo');
    modal.classList.remove('show');
    setTimeout(() => {
        modal.style.display = 'none';
        modalVideo.pause();
        modalVideo.src = '';
    }, 300);
}

function copyVideoUrl(videoUrl) {
    const fullUrl = window.location.origin + videoUrl;

    fs.copyFile(fullUrl, destinationPath, (err) => {
        if (err) {
            console.error('Dosya kopyalanırken hata oluştu:', err);
        } else {
            console.log('Dosya başarıyla kopyalandı:', destinationPath);
        }
    });


    navigator.clipboard.writeText(fullUrl)
        .then(() => {
            const button = event.target.closest('.action-button');
            const originalColor = button.style.color;
            button.style.color = '#22c55e'; // Success color
            setTimeout(() => {
                button.style.color = originalColor;
            }, 1000);
        })
        .catch(err => {
            console.error('Failed to copy URL:', err);
        });
}

// Close modal when clicking outside
window.onclick = function(event) {
    const modal = document.getElementById('videoModal');
    if (event.target == modal) {
        closeModal();
    }
}

// Close modal with escape key
document.addEventListener('keydown', function(event) {
    if (event.key === 'Escape') {
        closeModal();
    }
});

// Load videos when page loads
document.addEventListener('DOMContentLoaded', fetchVideos);

function toggleDownloadPopup() {
    const popup = document.getElementById('downloadPopup');
    popup.classList.toggle('show');
}

function validateInput(input) {
    const submitButton = document.getElementById('submitButton');
    submitButton.disabled = !input.checkValidity();
}

async function handleDownload(event) {
    event.preventDefault();
    const input = document.getElementById('videoUrl');
    const submitButton = document.getElementById('submitButton');
    const originalText = submitButton.textContent;
    
    try {
        submitButton.disabled = true;
        submitButton.textContent = 'İndiriliyor...';

        const response = await fetch('/dowland', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                videoString: input.value
            })
        });

        if (!response.ok) {
            throw new Error('Video indirilemedi');
        }

        // İndirme başarılı
        input.value = '';
        submitButton.textContent = 'Başarılı!';
        submitButton.style.background = '#22c55e';

        // Listeyi yenile
        setTimeout(() => {
            fetchVideos();
        }, 1000);

        // Butonu sıfırla
        setTimeout(() => {
            submitButton.textContent = originalText;
            submitButton.style.background = '';
            submitButton.disabled = false;
            toggleDownloadPopup();
        }, 2000);

    } catch (error) {
        console.error('Download error:', error);
        submitButton.textContent = 'Hata!';
        submitButton.style.background = '#ef4444';
        
        setTimeout(() => {
            submitButton.textContent = originalText;
            submitButton.style.background = '';
            submitButton.disabled = false;
        }, 2000);
    }
}

// Popup dışına tıklandığında kapat
document.addEventListener('click', function(event) {
    const popup = document.getElementById('downloadPopup');
    const downloadButton = event.target.closest('.download-button');
    const downloadPopup = event.target.closest('.download-popup');
    
    if (!downloadButton && !downloadPopup && popup.classList.contains('show')) {
        popup.classList.remove('show');
    }
});

async function deleteVideo(filename) {
    if (!confirm('Bu videoyu silmek istediğinizden emin misiniz?')) {
        return;
    }

    try {
        const response = await fetch(`/api/delete?filename=${filename}`, {
            method: 'DELETE'
        });

        if (!response.ok) {
            throw new Error('Video silinemedi');
        }

        // Silme başarılı, listeyi yenile
        await fetchVideos();

        // Başarı mesajı göster
        const toast = document.createElement('div');
        toast.className = 'fixed bottom-4 left-1/2 transform -translate-x-1/2 bg-green-500 text-white px-4 py-2 rounded-lg shadow-lg z-50';
        toast.textContent = 'Video başarıyla silindi';
        document.body.appendChild(toast);

        // 3 saniye sonra toast'ı kaldır
        setTimeout(() => {
            toast.remove();
        }, 3000);

    } catch (error) {
        console.error('Delete error:', error);
        alert('Video silinirken bir hata oluştu');
    }
}