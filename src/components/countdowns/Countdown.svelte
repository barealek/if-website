<script>
   export let date, text;

   import {onMount} from "svelte"

   let dateObj = new Date(date)

   onMount(() => {
      const interval = setInterval(() => {
         dateObj = new Date(dateObj)
      }, 500)

      return () => clearInterval(interval)
   })

   $: daysLeft = Math.floor((dateObj - new Date()) / 1000 / 60 / 60 / 24)
   $: hoursLeft = Math.floor((dateObj - new Date()) / 1000 / 60 / 60) % 24
   $: minutesLeft = Math.floor((dateObj - new Date()) / 1000 / 60) % 60
   $: secondsLeft = Math.floor((dateObj - new Date()) / 1000) % 60
</script>

<p class="text-white z-10 text-lg">
   <span class="font-bold tracking-wide">{text}</span> -
   <span class="text-clock tracking-tight">{daysLeft} {daysLeft != 1 ? "dage" : "dag"}, {hoursLeft} {hoursLeft != 1 ? "timer" : "time"}, {minutesLeft} {minutesLeft != 1 ? "minutter" : "minut"} og {secondsLeft} {secondsLeft != 1 ? "sekunder" : "sekund"}</span>
</p>
