function createCountdown(date, text) {
    const dateObj = new Date(date);
    const countdownElement = document.createElement('p');
    countdownElement.className = 'z-10 text-lg';

    function updateCountdown() {
        const now = new Date();
        const timeLeft = dateObj - now;
        if (timeLeft <= 0) {
            const timeSince = now - dateObj;
            const daysSince = Math.floor(timeSince / 1000 / 60 / 60 / 24);

            countdownElement.innerHTML = `
                <span class="font-bold tracking-wide">${text}</span> - <span class="text-clock tracking-tight">For ${daysSince} ${daysSince != 1 ? "dage" : "dag"} siden</span>
            `;
            clearInterval(interval);
            return;
        }
        const daysLeft = Math.floor(timeLeft / 1000 / 60 / 60 / 24);
        const hoursLeft = Math.floor(timeLeft / 1000 / 60 / 60) % 24;
        const minutesLeft = Math.floor(timeLeft / 1000 / 60) % 60;
        const secondsLeft = Math.floor(timeLeft / 1000) % 60;

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

    const interval = setInterval(updateCountdown, 500);
    updateCountdown();

    return countdownElement;
}

const countdownsContainer = document.getElementById('countdowns');
if (countdownsContainer) {
   const countdowns = {
      '2024-10-13 00:00': 'Efterårsferie',
      '2025-01-11 03:20': 'Fødselsdag',
      '2024-12-20 15:30': 'Juleferie',
   };

   for (const [date, text] of Object.entries(countdowns)) {
      const countdown = createCountdown(date, text);
      countdownsContainer.appendChild(countdown);
   }
}
