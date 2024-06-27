import "./Navbar.scss"

export default function Navbar() {
    return (
        <nav class={"huge-text"}>
            <p class="is-abril">
                Architects Land
            </p>
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
                    Équipes
                </a>
            </div>
        </nav>
    )
}