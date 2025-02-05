<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>countdowns</title>
   <link rel="stylesheet" href="/css/compiled.css">
</head>
<body>

   <?php include $_SERVER['DOCUMENT_ROOT'] . '/lib/nav.php'; ?>

   <!-- Card Grid -->
   <div class="grid grid-cols-2 gap-4 py-12 sm:gap-6 md:gap-8 lg:gap-12">
   <!-- Card -->
      <a class="flex flex-col group focus:outline-none" href="#">
         <div class="overflow-hidden bg-gray-100 aspect-w-16 aspect-h-12 rounded-2xl">
            <img src="/img/jackie.webp" alt="Blog Image" class="object-cover transition-transform duration-500 ease-in-out group-hover:scale-105 group-focus:scale-105 rounded-2xl">
         </div>
      </a>
      <a class="flex flex-col group focus:outline-none" href="#">
         <div class="overflow-hidden bg-gray-100 aspect-w-16 aspect-h-12 rounded-2xl">
            <img src="/img/stresa.webp" alt="Blog Image" class="object-cover transition-transform duration-500 ease-in-out group-hover:scale-105 group-focus:scale-105 rounded-2xl">
         </div>
      </a>
   <!-- End Card -->
   <div class="px-12">
      <h2 class="text-xl font-semibold">Mathgamers video</h2>
      <iframe width="560" height="315" src="https://www.youtube.com/embed/10l8Doumb9o?si=5EoO8rekPVUsdDMr" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
   </div>
   </div>
<!-- End Card Grid -->
</body>
</html>
