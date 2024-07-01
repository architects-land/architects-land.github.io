import Root from "./pages/root/Root.tsx";
import Navbar from "./components/navbar/Navbar.tsx";
import Lost from "./pages/lost/Lost.tsx";
import Footer from "./components/footer/Footer.tsx";
import Rules from "./pages/rules/Rules.tsx";
import Team from "./pages/team/Team.tsx";
import Season from "./pages/season/Season.tsx";

function App() {
  const href = location.pathname;

  switch (href) {
    case "/":
      return (
        <>
          <Navbar />
          <Root />
          <Footer />
        </>
      );
    case "/rules":
      document.title = "Règles - Architects Land";
      return (
        <>
          <Navbar />
          <Rules />
          <Footer />
        </>
      );
    case "/team":
      document.title = "Équipe - Architects Land";
      return (
        <>
          <Navbar />
          <Team />
          <Footer />
        </>
      );
  }
  if (href.startsWith("/season/")) {
    const id = href.split("/")[2];
    return (
      <>
        <Navbar />
        <Season id={id} />
        <Footer />
      </>
    );
  }
  setTimeout((_: any) => {
    location.replace(`${location.origin}`);
  }, 5 * 1000);
  return (
    <>
      <Lost />
    </>
  );
}

export default App;

export function setupEvents() {
  const clickables: HTMLCollectionOf<Element> =
    document.getElementsByClassName("is-clickable");
  const nav = document.getElementById("navbar")!!;

  let blocked: Element[] = [];

  if (window.innerWidth > 620) {
    nav.style.display = "none";
  }

  for (const c of clickables) {
    const el = c as HTMLElement;
    const href = c.getAttribute("data-href")!!;
    if (href == "") el.style.cursor = "default";
    c.addEventListener("click", (_) => {
      if (href.startsWith("http")) {
        location.href = href;
      } else if (href != "") {
        location.href = `${location.origin}${href}`;
      }
    });
    c.addEventListener("mouseenter", (_) => {
      if (!blocked.includes(c)) {
        c.classList.add("is-clickable--active");
        addToBlocked(c);
        return;
      }
      updateStatus(c, "is-clickable");
    });
    c.addEventListener("mouseleave", (_) => {
      if (!blocked.includes(c)) {
        c.classList.remove("is-clickable--active");
        addToBlocked(c);
        return;
      }
      updateStatus(c, "is-clickable", false);
    });
  }

  if (window.innerWidth > 620) {
    document.addEventListener("scroll", (_) => {
      if (window.scrollY >= window.innerHeight) {
        nav.style.display = "";
        nav.classList.add("nav--present");
        return;
      }
      nav.classList.remove("nav--present");
    });
  }

  function updateStatus(el: Element, baseClass: string, add = true) {
    setTimeout((_: any) => {
      if (!blocked.includes(el)) {
        if (add) el.classList.add(`${baseClass}--active`);
        else el.classList.remove(`${baseClass}--active`);
        addToBlocked(el);
        return;
      }
      updateStatus(el, baseClass, add);
    }, 10);
  }

  function addToBlocked(el: Element) {
    blocked.push(el);
    setTimeout((_: any) => {
      blocked.splice(blocked.indexOf(el), 1);
    }, 1000);
  }
}
