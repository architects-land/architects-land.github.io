import Hero from "../../components/hero/Hero.tsx";
import "./Team.scss";

export default function Team() {
  return (
    <>
      <Hero
        image={"https://loremflickr.com/1920/1080"}
        title={"Équipe"}
        description={""}
        min
      />
      <main>
        <div class="content team__presentation large-text">
          <h2>Les créateurs d'Architects Land</h2>
          <p>
            Architects Land, c'est aussi une équipe dévouée visant à créer un
            SMP inoubliable !
          </p>
          <p>
            Architects Land n'est pas qu'un serveur Minecraft, c'est aussi des
            modpacks avec des mods customs, des événements nécessitant des
            plugins créés de toute pièce et une site web pour présenter le
            projet.
          </p>
        </div>
        <div class="content team__members large-text">
          <Member
            name={"Anhgelus Morhtuuzh"}
            image={"/skins/anhgelus.png"}
            description={
              "Fondateur et développeur, Anhgelus Morhtuuzh gère le serveur et prend toutes les décisions importantes."
            }
          />
          <Member
            name={"Akik4"}
            image={"/skins/Akik4.png"}
            description={
              "Akik4 est un développeur Spigot aidant grandement Anhgelus Morhtuuzh."
            }
          />
        </div>
      </main>
    </>
  );
}

function Member(props: any) {
  return (
    <div class={"team__member is-clickable"} data-href={props.link}>
      <img src={props.image} alt={`Skin de ${props.name}`} />
      <h5>{props.name}</h5>
      <p>{props.description}</p>
    </div>
  );
}
