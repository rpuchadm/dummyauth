import React from 'react'

import Auth from './auth'
import HeaderMenu from './headermenu'
import Layout from './layout'


const Root = () => {

    return (

        <Layout>
            <HeaderMenu />
            <Auth />
        </Layout>

    )
}

export default Root