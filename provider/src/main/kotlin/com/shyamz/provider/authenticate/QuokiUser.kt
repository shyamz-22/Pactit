package com.shyamz.provider.authenticate

import javax.persistence.*

@Entity
data class QuokiUser(
        @Id
        @GeneratedValue(strategy = GenerationType.IDENTITY)
        val id: Int = 0,
        val userName: String,
        val password: String)