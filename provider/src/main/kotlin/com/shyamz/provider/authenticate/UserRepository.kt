package com.shyamz.provider.authenticate

import org.springframework.data.repository.CrudRepository
import org.springframework.stereotype.Repository

@Repository
interface UserRepository: CrudRepository<QuokiUser, Long> {
    fun findByUserName(userName: String): QuokiUser?
}