import "./Footer.scss";

export default function Footer() {
  return (
    <footer class={"footer"}>
      <img src="/logo-sand.png" alt="Logo" class={"footer__logo"} />
      <div class="footer__links large-text">
        <a href="/">Accueil</a>
        <a href="/#seasons">Saisons</a>
        <a href="/rules">Règles</a>
        <a href="/team">Équipe</a>
        <a href="https://www.anhgelus.world/#legal-fr" target={"_blank"}>
          Mentions Légales
        </a>
      </div>
      <div class={"footer__credits"}>
        <div class={"footer__links footer__links--inline"}>
          <a href={"https://youtu.be/_SP_1hrsZBU"} target={"_blank"}>
            Présentation
          </a>
          <a href={"https://github.com/architects-land"} target={"_blank"}>
            GitHub
          </a>
        </div>
        <br />
        <p>Contenu CC BY-SA-NC &copy; 2024 Architects Land.</p>
        <p>Code source AGPL-3.0 &copy; 2024 Architects Land.</p>
        <br />
        <p>
          Site web créé et maintenu par{" "}
          <a href="https://www.anhgelus.world" target={"_blank"}>
            William Hergès
          </a>
          .
        </p>
        <p>Créé à l'aide de SolidJS, Vite, SCSS et Bun !</p>
      </div>
    </footer>
  );
}
