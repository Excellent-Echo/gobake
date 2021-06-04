import React from 'react'
import "./card.scss";
import products from '../../../dataProduct/data';
import {numberWithCommas} from '../../../utils/utils';

function Card() {
    return(
        products.map((product) => {
            return (
                <div key={product.id} className="card-wrapper">
                    <div className="card">
                        <div className="card-image">
                            <img src={product.product_image} alt={product.product_name} />
                        </div>
                        <div className="card-data">
                            <h3 className="title-card">
                                {product.product_name}
                            </h3>
                            <span className="price-card">
                                {`Rp ${numberWithCommas(product.price)}`}
                            </span>
                        </div>
                    </div>
                </div>
            )
        })
    )
}

export default Card
