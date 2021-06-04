import React from 'react'
import { banner } from '../../../assets';
import "./banner.scss";

function Banner() {
    return (
        <div className="banner-wrapper">
            <div className="banner">
                <div className="banner-text">
                    <h3>Good for your tummy</h3>
                    <h1 className="title-banner">Gen-XYZ Favourite Bakery</h1>
                    <button className="order-btn">Pesan</button>
                </div>
                <div className="banner-image">
                    <img src={banner} alt="banner" />
                </div>
            </div>
        </div>
    )
}

export default Banner
