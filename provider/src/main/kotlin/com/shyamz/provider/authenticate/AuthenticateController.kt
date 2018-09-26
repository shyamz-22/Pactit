package com.shyamz.provider.authenticate

import org.springframework.http.ResponseEntity
import org.springframework.util.MultiValueMap
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/users/{username}/authentication")
class AuthenticateController(private val userRepository: UserRepository) {

    @PostMapping
    fun authenticate(@PathVariable username: String,
                     @RequestParam requestParams: MultiValueMap<String, String> ): ResponseEntity<Unit> {

        return userRepository.findByUserName(username)
                ?.takeIf { it.password == requestParams.getFirst("password") }
                ?.let { ResponseEntity.noContent().build<Unit>() } ?: ResponseEntity.notFound().build()

    }
}