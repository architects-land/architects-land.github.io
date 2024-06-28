import "./Navbar.scss"

export default function Navbar() {
    return (
        <nav class={"huge-text nav nav--hidden"} id="navbar" role="navigation">
            <a class="is-abril" href={"/"}>
                Architects Land
            </a>
            <div class="nav__links">
                <a href={"/"}>
                    Accueil
                </a>
                <a href={"/#seasons"}>
                    Saisons
                </a>
                <a href={"/rules"}>
                    Règles
                </a>
                <a href={"/team"}>
                    Équipe
                </a>
            </div>
        </nav>
    )
}