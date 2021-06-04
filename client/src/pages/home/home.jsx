import React from "react";
import { Banner, Card } from "../../components/molecules";
import "./home.scss";

function Home() {
  return (
    <div className="home-wrapper">
      <div className="home">
        <div className="main-banner">
          <Banner /> 
        </div>
        <div className="main-card">
          <div className="main-card-title"><h2>Menu Pilihan</h2></div>
          <div className="cards">
            <Card />
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;
