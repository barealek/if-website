<!DOCTYPE html>
<html lang="en">
<head>
   <meta charset="UTF-8">
   <meta name="viewport" content="width=device-width, initial-scale=1.0">
   <title>Document</title>
   <link rel="stylesheet" href="/css/compiled.css">
   <script defer src="/js/navbar.js"></script>
</head>
<body>
   <div id="navbar"></div>

   <section class="flex flex-col justify-center items-center h-[75vh] w-screen">

      <img src="/img/alek.jpg" alt="billed af mig" class="object-cover rounded-full aspect-square w-72 ring-2 ring-offset-1 ring-blue-200">
      <?php
      $bog = [
          ["Navn", "Alek"],
          ["Øjenfarve", "Blå"],
          ["Højde", "193 cm"],
          ["Kaldenavn", "Aleksand eller falsk"],
          ["Fejlede startups", "12"],
          ["Konto", "60 kr."],
          ["Interesser", "Kode, klatre"]
      ];
      ?>

      <div class="flex items-center justify-center mt-8">
          <div class="grid text-lg text-neutral-900 bg-neutral-100">
              <?php foreach ($bog as $i => $entry): ?>
                  <div class="flex gap-4 <?= $i === 0 ? '' : 'border-t border-purple-500' ?> py-1">
                      <p class="font-bold"><?= htmlspecialchars($entry[0]) ?></p>
                      <p class="w-full text-right"><?= htmlspecialchars($entry[1]) ?></p>
                  </div>
              <?php endforeach; ?>
          </div>
      </div>

   </section>
</body>
</html>
