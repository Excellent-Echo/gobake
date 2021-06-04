import React from 'react'
import './menu.scss'
import {Card} from '../../components/index'

function Menu() {
    return (
        <div className="menu-wrapper">
            <div className="menu">
                <div className="menu-title">
                    <h2>Menus</h2>
                </div>
                <div className="menu-card">
                    <Card />
                </div>
            </div>
        </div>
    )
}

export default Menu
