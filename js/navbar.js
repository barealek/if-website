const tmpl = `
<nav class="bg-white shadow">
   <div class="container flex items-center justify-center p-6 mx-auto text-gray-600 capitalize">
      <a href="/">Blå bog</a>

      <a href="/countdowns/">countdowns</a>

      <a href="/studietur/">studietur</a>

      <a href="/medie/">medie</a>

      <a href="https://maps.app.goo.gl/KnnhpjvoVXTJZFqS6" target="_blank">kort</a>
   </div>
</nav>
`

elm = document.getElementById('navbar')
elm.innerHTML = tmpl

const unselectedStyles = "border-b-2 border-transparent hover:text-gray-800 transition-colors duration-300 transform hover:border-blue-500 mx-1.5 sm:mx-6"
const selectedStyles = "text-gray-800 transition-colors duration-300 transform border-b-2 border-blue-500 mx-1.5 sm:mx-6"

const links = document.querySelectorAll('nav a')
links.forEach(link => {
   if (link.href === window.location.href) {
      link.classList = selectedStyles
   } else {
      link.classList = unselectedStyles
   }
})
