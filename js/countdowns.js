function createCountdown(date, text) {
    const dateObj = new Date(date);
    const countdownElement = document.createElement('p');
    countdownElement.className = 'z-10 text-lg';

    function updateCountdown() {
        const now = new Date();
        const daysLeft = Math.floor((dateObj - now) / 1000 / 60 / 60 / 24);
        const hoursLeft = Math.floor((dateObj - now) / 1000 / 60 / 60) % 24;
        const minutesLeft = Math.floor((dateObj - now) / 1000 / 60) % 60;
        const secondsLeft = Math.floor((dateObj - now) / 1000) % 60;

        countdownElement.innerHTML = `
            <span class="font-bold tracking-wide">${text}</span> -
            <span class="text-clock tracking-tight">
                ${daysLeft} ${daysLeft != 1 ? "dage" : "dag"},
                ${hoursLeft} ${hoursLeft != 1 ? "timer" : "time"},
                ${minutesLeft} ${minutesLeft != 1 ? "minutter" : "minut"} og
                ${secondsLeft} ${secondsLeft != 1 ? "sekunder" : "sekund"}
            </span>
        `;
    }

    setInterval(updateCountdown, 500);
    updateCountdown();

    return countdownElement;
}

const countdownsContainer = document.getElementById('countdowns');
if (countdownsContainer) {
   const countdowns = {
      '2024-12-31 23:59': 'Nytår',
      '2025-01-11 03:20': 'Fødselsdag'
   };

   for (const [date, text] of Object.entries(countdowns)) {
      const countdown = createCountdown(date, text);
      countdownsContainer.appendChild(countdown);
   }
}
