
<nav class="bg-white shadow">
   <?php
   $unselectedStyles = "border-b-2 border-transparent hover:text-gray-800 transition-colors duration-300 transform hover:border-blue-500 mx-1.5 sm:mx-6";
   $selectedStyles = "text-gray-800 transition-colors duration-300 transform border-b-2 border-blue-500 mx-1.5 sm:mx-6";

   $links = [
      "/" => "Blå bog",
      "/countdowns/" => "countdowns",
      "/studietur/" => "studietur",
      "/medie/" => "medie",
      "https://maps.app.goo.gl/KnnhpjvoVXTJZFqS6" => "kort",
      "/login" => "login",
      "/lorem" => "lorem"
   ];

   $current = $_SERVER['REQUEST_URI'];
   ?>
   <div class="container flex items-center justify-center p-6 mx-auto text-gray-600 capitalize">
      <?php foreach ($links as $url => $label): ?>
         <?php $styles = (strpos($current, $url) !== false) ? $selectedStyles : $unselectedStyles; ?>
         <a href="<?= $url ?>" class="<?= $styles ?>"><?= $label ?></a>
      <?php endforeach; ?>
   </div>
</nav>
