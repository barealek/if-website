<script>
	import bog from "../data/bog.json"
</script>

<div class="flex items-center justify-center mt-8">
	<div class="grid text-lg text-neutral-100">
		{#each bog as entry, i}
			<div class="flex gap-4 {i === 0 ? '' : 'border-t border-purple-500'} py-1 ">
				<p class="font-bold">{entry[0]}</p>
				<p class="text-right w-full">{entry[1]}</p>
			</div>
		{/each}
	</div>
</div>