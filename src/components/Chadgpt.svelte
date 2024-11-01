<script>
   import { onMount } from 'svelte'
   import ChadgptTitle from './ChadgptTitle.svelte'

   let isLoaded = false
   let state = "waiting:user"

   let messages = []

   let hintMessage

   let ws

   onMount(() => {
      // Connect to the websocket at /api/chadgpt
      ws = new WebSocket('/api/chad')
      ws.onopen = () => {
         console.log('Connected to the websocket')
      }
      ws.onmessage = (event) => {
         const data = JSON.parse(event.data)
         if (data.system == "start") {
            isLoaded = true
         } else if (data.system == "update") {
            messages = data.messages
         } else if (data.system == "state") {
            state = data.state
         } else if (data.system == "hint") {
            hintMessage = data.hint
         }
      }
      ws.onclose = () => {
         console.log('Disconnected from the websocket')
         isLoaded = false
         alert("Disconnected fra websocketen")
      }
   })


   let message = ""
   function submitMessage() {
      console.log("Sender " + message)
      if (message.trim() === "") return
      if (state !== "waiting:user") return
      ws.send(message)
      message = ""
   }

</script>

{#if isLoaded}
<!-- Content -->
<div class="relative h-screen">
  <div class="py-10 lg:py-14">
    <!-- Title -->
    <ChadgptTitle />
    state: {state}
    <!-- End Title -->

    <ul class="mt-16 space-y-5">
      <!-- Chat Bubble -->
      <li class="max-w-4xl py-2 px-4 sm:px-6 lg:px-8 mx-auto flex gap-x-2 sm:gap-x-4">
      <img src="/img/chadden.jpg" alt="chadden" class="shrink-0 size-[38px] rounded-full">

        <div class="space-y-3">
          <h2 class="font-medium text-gray-800 dark:text-white">
            Hvad kan jeg hjælpe med?
          </h2>
          <div class="space-y-1.5">
            <p class="mb-1.5 text-sm text-gray-800 dark:text-white">
               Du kan spørge mig om ting såsom:
            </p>
            <ul class="list-disc list-outside space-y-1.5 ps-3.5">
              <li class="text-sm text-gray-800 dark:text-white">
                Hvad er klokken?
              </li>

              <li class="text-sm text-gray-800 dark:text-white">
                  Hvad dag er det i dag?
              </li>

              <li class="text-sm text-gray-800 dark:text-white">
               Hvordan er vejret i ballerup?
              </li>
            </ul>
          </div>
        </div>
      </li>
      <!-- End Chat Bubble -->

      <!-- Chat Bubble -->
      {#each messages as msg}
      <li class="py-2 sm:py-4">
        <div class="max-w-4xl px-4 sm:px-6 lg:px-8 mx-auto">
          <div class="max-w-2xl flex gap-x-2 sm:gap-x-4">
            <span class="shrink-0 inline-flex items-center justify-center size-[38px] rounded-full bg-gray-600">
               {#if msg.author === 'chadgpt'}
                  <img src="/img/chadden.jpg" alt="chadden" class="shrink-0 size-[38px] rounded-full">
               {:else}
                  <span class="text-sm font-medium text-white leading-none">dig</span>
               {/if}
            </span>

            <div class="grow mt-2 space-y-3">
              <p class="text-gray-800 dark:text-neutral-200">
                {msg.content}
              </p>
            </div>
          </div>
        </div>
      </li>
      {/each}
      <!-- End Chat Bubble -->

      {#if state === "waiting:chad"}
      <li class="py-2 sm:py-4">
        <div class="max-w-4xl px-4 sm:px-6 lg:px-8 mx-auto">
          <div class="max-w-2xl flex gap-x-2 sm:gap-x-4">
            <span class="shrink-0 inline-flex items-center justify-center size-[38px] rounded-full bg-gray-600">
               <img src="/img/chadden.jpg" alt="chadden" class="shrink-0 size-[38px] rounded-full">
            </span>

            <div class="grow mt-2 space-y-3">
              <p class="text-gray-800 dark:text-neutral-200">
                {hintMessage === "hint::" ? "Tænker..." : hintMessage.replace("hint::", "")}
              </p>
            </div>
          </div>
        </div>
      </li>
      {/if}


    </ul>
  </div>

  <div class="sticky bottom-0 z-10 bg-white border-t border-gray-200 pt-2 pb-3 sm:pt-4 sm:pb-6 dark:bg-neutral-900 dark:border-neutral-700">
    <!-- Textarea -->
    <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">

      <!-- Input -->
      <div class="relative">
        <textarea bind:value={message} class="p-4 pb-12 block w-full border-gray-200 rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none dark:bg-neutral-900 dark:border-neutral-700 dark:text-neutral-400 dark:placeholder-neutral-500 dark:focus:ring-neutral-600" placeholder="Spørg løs..."></textarea>

        <!-- Toolbar -->
        <div class="absolute bottom-px inset-x-px p-2 rounded-b-lg bg-white dark:bg-neutral-900">
          <div class="flex justify-between items-center">
            <!-- Button Group -->
            <div class="flex items-center">
              <!-- Mic Button -->
              <button type="button" class="inline-flex shrink-0 justify-center items-center size-8 rounded-lg text-gray-500 hover:bg-gray-100 focus:z-10 focus:outline-none focus:bg-gray-100 dark:text-neutral-500 dark:hover:bg-neutral-700 dark:focus:bg-neutral-700">
                <svg class="shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="3" rx="2"/><line x1="9" x2="15" y1="15" y2="9"/></svg>
              </button>
              <!-- End Mic Button -->

              <!-- Attach Button -->
              <button type="button" class="inline-flex shrink-0 justify-center items-center size-8 rounded-lg text-gray-500 hover:bg-gray-100 focus:z-10 focus:outline-none focus:bg-gray-100 dark:text-neutral-500 dark:hover:bg-neutral-700 dark:focus:bg-neutral-700">
                <svg class="shrink-0 size-4" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m21.44 11.05-9.19 9.19a6 6 0 0 1-8.49-8.49l8.57-8.57A4 4 0 1 1 18 8.84l-8.59 8.57a2 2 0 0 1-2.83-2.83l8.49-8.48"/></svg>
              </button>
              <!-- End Attach Button -->
            </div>
            <!-- End Button Group -->

            <!-- Button Group -->
            <div class="flex items-center gap-x-1">

              <!-- Send Button -->
              <button on:click={submitMessage} type="button" class="inline-flex shrink-0 justify-center items-center size-8 rounded-lg text-white bg-blue-600 hover:bg-blue-500 focus:z-10 focus:outline-none focus:bg-blue-500">
                <svg class="shrink-0 size-3.5" xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" viewBox="0 0 16 16">
                  <path d="M15.964.686a.5.5 0 0 0-.65-.65L.767 5.855H.766l-.452.18a.5.5 0 0 0-.082.887l.41.26.001.002 4.995 3.178 3.178 4.995.002.002.26.41a.5.5 0 0 0 .886-.083l6-15Zm-1.833 1.89L6.637 10.07l-.215-.338a.5.5 0 0 0-.154-.154l-.338-.215 7.494-7.494 1.178-.471-.47 1.178Z"/>
                </svg>
              </button>
              <!-- End Send Button -->
            </div>
            <!-- End Button Group -->
          </div>
        </div>
        <!-- End Toolbar -->
      </div>
      <!-- End Input -->
    </div>
    <!-- End Textarea -->
  </div>
</div>
<!-- End Content -->

{/if}
