import "./People.scss";

export default function Card(props: any) {
  return (
    <div class={"people__card is-clickable"} data-href={props.link}>
      <div class="is-clickable__animation"></div>
      <img src={props.image} alt={`Skin de ${props.name}`} />
      <h5>{props.name}</h5>
      <p>{props.description}</p>
    </div>
  );
}
