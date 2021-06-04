import { FontAwesomeIcon, faInstagram, faFacebook, faTwitter } from "../../../assets/index";
import React from "react";
import "./footer.scss";

function Footer() {
  return (
    <div className="footer-wrapper">
     <div className="footer">
       <div className="footer-left">
          <h2 className="logo-footer">Gobake</h2>
       </div>
       <div className="footer-center">
          <span className="address-footer">Jl. Radio dalam Jakarta, Indonesia</span>
          <span className="phone-footer">+621 123456</span>
       </div>
       <div className="footer-right">
          <i className="footer-media"><FontAwesomeIcon icon={faTwitter} /></i>
          <i className="footer-media"><FontAwesomeIcon icon={faInstagram} /></i>
          <i className="footer-media"><FontAwesomeIcon icon={faFacebook} /></i>
       </div>
     </div>
    </div>
  );
}

export default Footer;
