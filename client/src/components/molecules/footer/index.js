
import React from "react";
import "./Footer.css";

function Footer() {
  return (
    <div className="main-footer">
      <div className="container">
        <div className="row">
          {/* Column1 */}
          <div className="col">
            <h4>GO-BAKE</h4>
            <h1 className="list-unstyled">
              <li>021-556-3737</li>
              <li>Jakarta, Indonesia</li>
              <li>Jl. Kaliadem No. 66 Tebet, jakarta Selatan</li>
            </h1>
          </div>
          {/* Column2 */}
          <div className="col">
            <h4>Sosial Media</h4>
            <ui className="list-unstyled">
              <li>Instagram</li>
              <li>Facebook</li>
              <li>Twitter</li>
            </ui>
          </div>
          {/* Column3 */}
          <div className="col">
            <h4>Our Service</h4>
            <ui className="list-unstyled">
              <li>Go-Bake Instan</li>
              <li>Go-Bake Delivery</li>
            </ui>
          </div>
        </div>
        <hr />
        <div className="row">
          <p className="col-sm">
            &copy;{new Date().getFullYear()} Go-Bake Indonesia | All rights reserved |
            Terms Of Service | Privacy
          </p>
        </div>
      </div>
    </div>
  );
}

export default Footer;