const elm = document.getElementById('countdowns')

const countdowns = [
   {
      id: "!",
      title: "Juleaften",
      date: "24 dec 2024"
   }
]

const tmpl = `
<div class="container mx-auto">
   <h1 class="text-3xl font-bold text-center">Countdowns</h1>
   <div class="grid grid-cols-1 gap-4 mt-6">
      ${countdowns.map(countdown => `
         <div class="bg-white shadow rounded-lg p-6">
            <h2 class="text-xl font-bold">${countdown.title}</h2>
            <p class="text-gray-600">${countdown.date}</p>
         </div>
      `).join('')}
   </div>
</div>
`

elm.innerHTML = tmpl
