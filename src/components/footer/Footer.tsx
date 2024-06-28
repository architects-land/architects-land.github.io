import "./Footer.scss"

export default function Footer() {
    return (
        <footer class={"footer"}>
            <h3>Architects Land</h3>
            <div class="footer__links large-text">
                <a href="/">Accueil</a>
                <a href="/#seasons">Saisons</a>
                <a href="/rules">Règles</a>
                <a href="/team">Équipe</a>
                <a href="https://www.anhgelus.world/#legal-fr" target={"_blank"}>Mentions Légales</a>
            </div>
            <div class={"footer__credits"}>
                <p>&copy; 2024 William Hergès</p>
                <p>Contenu sous licence CC BY-SA-NC</p>
                <p>Code source sous licence AGPL-3.0</p>
            </div>
        </footer>
    )
}