const tmpl = `
<nav class="bg-white shadow">
   <div class="container flex items-center justify-center p-6 mx-auto text-gray-600 capitalize">
      <a href="/">Blå bog</a>

      <a href="/countdowns/">countdowns</a>

      <a href="/studietur/">studietur</a>

      <a href="/kort/">kort</a>
   </div>
</nav>
`

elm = document.getElementById('navbar')
elm.innerHTML = tmpl

const selectedStyles = "border-b-2 border-transparent hover:text-gray-800 transition-colors duration-300 transform hover:border-blue-500 mx-1.5 sm:mx-6"
const unselectedStyles = "text-gray-800 transition-colors duration-300 transform border-b-2 border-blue-500 mx-1.5 sm:mx-6"

const links = document.querySelectorAll('nav a')
links.forEach(link => {
   if (link.href === window.location.href) {
      link.classList = unselectedStyles
   } else {
      link.classList = selectedStyles
   }
})
