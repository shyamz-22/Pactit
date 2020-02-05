package com.shyamz.provider.authenticate

import org.springframework.beans.factory.annotation.Value
import org.springframework.http.HttpStatus
import org.springframework.http.MediaType
import org.springframework.http.ResponseEntity
import org.springframework.util.MultiValueMap
import org.springframework.web.bind.annotation.*

@RestController
@RequestMapping("/users/{username}/authentication")
class AuthenticateController(private val userRepository: UserRepository,
                             @Value("\${api.key}") private val apiKey: String) {

    @PostMapping
    fun authenticate(@PathVariable username: String,
                     @RequestParam requestParams: MultiValueMap<String, String>,
                     @RequestHeader("X-Api-Key") apiKey: String?): ResponseEntity<AuthenticateResult> {

        if (!validKey(apiKey)) {
            return ResponseEntity
                    .status(HttpStatus.UNAUTHORIZED)
                    .build()
        }

        val result = userRepository.findByUserName(username)
                ?.takeIf { it.password == requestParams.getFirst("password") }
                ?.let { AuthenticateResult.Success(true) } ?: AuthenticateResult.Error("Invalid username and password")

        return when(result) {
            is AuthenticateResult.Success -> ResponseEntity.status(HttpStatus.NO_CONTENT).build()
            is AuthenticateResult.Error -> ResponseEntity
                    .badRequest()
                    .contentType(MediaType.APPLICATION_JSON)
                    .body(result)
        }

    }

    private fun validKey(apiKey: String?): Boolean {
        return this.apiKey == apiKey
    }
}