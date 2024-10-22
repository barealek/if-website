<script>
   import { onMount } from "svelte";

   let history = ''
   onMount(() => {
      var k = window.addEventListener('keydown', (e) => {
         history = history + e.key.toLowerCase()

         if (history.includes('maccen')) {
            visible = true
            history = ''
         }

      })

      return () => {
         window.removeEventListener('keydown', k)
      }
   })

   let visible = false

</script>

{#if visible}
<div class="absolute inset-0 bg-black/25 h-screen w-screen flex items-center justify-center">

<section class="max-w-4xl p-6 mx-auto bg-white rounded-md shadow-md">
   <div class="flex w-full justify-between">
      <h2 class="text-lg font-semibold text-gray-700 capitalize">Find vej til maccen</h2>
      <button on:click={() => {visible = false}} class="p-2 text-gray-700 text-xl">x</button>

   </div>

    <iframe title="ballerup maccen" src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d71859.64135149382!2d12.29502217735326!3d55.74783907896231!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x4652507a5d1a12ed%3A0xd9965114c4f760c4!2sMcDonald&#39;s%20Ballerup!5e0!3m2!1sda!2sse!4v1729601485605!5m2!1sda!2sse" width="600" height="450" style="border:0;" allowfullscreen="" loading="lazy" referrerpolicy="no-referrer-when-downgrade"></iframe>
</section>
</div>

{/if}
