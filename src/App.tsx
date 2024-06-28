import Root from "./pages/root/Root.tsx";
import Navbar from "./components/navbar/Navbar.tsx";
import Lost from "./pages/lost/Lost.tsx";
import Footer from "./components/footer/Footer.tsx";
import Rules from "./pages/rules/Rules.tsx";

function App() {
    const href = location.pathname

    switch (href) {
        case "/":
            return (
                <>
                    <Navbar />
                    <Root />
                    <Footer />
                </>
            )
        case "/rules":
            return (
                <>
                    <Navbar />
                    <Rules />
                    <Footer />
                </>
            )
    }
    setTimeout((_: any) => {
        location.replace(`${location.origin}`)
    }, 5*1000)
    return (
        <>
            <Lost />
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
            location.replace(`${location.origin}/season/${href}`)
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
