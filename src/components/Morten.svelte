<script lang="ts">
	import ArcRotateCamera from 'svelte-babylon/components/Cameras/ArcRotateCamera/index.svelte'
	import Canvas from 'svelte-babylon/components/Canvas/index.svelte'
	import DirectionalLight from 'svelte-babylon/components/Lights/DirectionalLight/index.svelte'
	import HemisphericLight from 'svelte-babylon/components/Lights/HemisphericLight/index.svelte'
	import Custom from 'svelte-babylon/components/Objects/Custom/index.svelte'
	import Ground from 'svelte-babylon/components/Objects/Ground/index.svelte'
	import Scene from 'svelte-babylon/components/Scene/index.svelte'
	import { Vector3 } from '@babylonjs/core/Maths/math.vector'
	import type { Mesh } from '@babylonjs/core/Meshes/mesh.js'
	import type { Writable } from 'svelte/types/runtime/store'

	let model: Writable<Mesh>
	let shadowObjects: Array<Mesh>
	$: if (model) {
		shadowObjects = [$model]
	}

   let rotationDeg = 0

   setInterval(() => {
      rotationDeg = (rotationDeg + 0.01)
   }, 30);
</script>

<Canvas
	antialiasing={true}
	engineOptions={{
		preserveDrawingBuffer: true,
		stencil: true,
	}}
>
	<Scene>
		<HemisphericLight intensity={0.5} />
		<DirectionalLight
			intensity={0.25}
			direction={new Vector3(-10, -20, -10)}
			position={new Vector3(2, 6, 2)}
			castShadowOf={shadowObjects}
		/>
		<ArcRotateCamera target={new Vector3(0, 2.7, 0)} disableControl={true} />
		<Custom
			url="/models/23.10.2024.glb"
			scaling={new Vector3(1.5, 1.5, 1.5)}
			position={new Vector3(0, 0.55, 0)}
			receiveShadows
         rotation={new Vector3(0, rotationDeg, 0)}
         bind:object={model}
		/>
	</Scene>
</Canvas>
