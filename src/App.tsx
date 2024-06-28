import Root from "./pages/root/Root.tsx";
import Navbar from "./components/navbar/Navbar.tsx";

function App() {
  return (
    <>
        <Navbar />
      <Root />
    </>
  )
}

export default App

export function setupEvents() {
    const seasons = document.getElementsByClassName('presentation__season')
    const nav = document.getElementById("navbar")!!

    let blocked: Element[] = []

    nav.style.display = 'none'

    for (const season of seasons) {
        season.addEventListener('mouseenter', _ => {
            if (!blocked.includes(season)) {
                season.classList.add("presentation__season--active")
                addToBlocked(season)
                return;
            }
            updateStatus(season, "presentation__season")
        })
        season.addEventListener('mouseleave', _ => {
            if (!blocked.includes(season)) {
                season.classList.remove("presentation__season--active")
                addToBlocked(season)
                return;
            }
            updateStatus(season, "presentation__season", false)
        })
        season.addEventListener("click", _ => {
            const href = season.getAttribute("data-href")!!
            window.location.href = `/season/${href}`
        })
    }

    document.addEventListener("scroll", _ => {
        if (window.scrollY >= window.innerHeight) {
            nav.style.display = ''
            nav.classList.add("nav--present")
            return
        }
        nav.classList.remove("nav--present")
    })

    function updateStatus(el: Element, baseClass: string, add = true) {
        setTimeout((_: any) => {
            if (!blocked.includes(el)) {
                if (add) el.classList.add(`${baseClass}--active`)
                else el.classList.remove(`${baseClass}--active`)
                addToBlocked(el)
                return;
            }
            updateStatus(el, baseClass, add)
        }, 10)
    }

    function addToBlocked(el: Element) {
        blocked.push(el)
        setTimeout((_: any) => {
            blocked.splice(blocked.indexOf(el), 1)
        }, 1000);
    }
}
