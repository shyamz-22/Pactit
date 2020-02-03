package com.shyamz.provider.authenticate

import au.com.dius.pact.provider.junit.Consumer
import au.com.dius.pact.provider.junit.Provider
import au.com.dius.pact.provider.junit.State
import au.com.dius.pact.provider.junit.loader.PactFolder
import au.com.dius.pact.provider.junit.target.HttpTarget
import au.com.dius.pact.provider.junit.target.Target
import au.com.dius.pact.provider.junit.target.TestTarget
import au.com.dius.pact.provider.spring.SpringRestPactRunner
import org.junit.Before
import org.junit.runner.RunWith
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest


@RunWith(SpringRestPactRunner::class)
@Provider("UserManager")
@Consumer("Quoki")
@PactFolder("../consumer/src/consumer/http/pacts")
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.DEFINED_PORT,
        properties = ["server.port=8601"])
class QuokiUserAuthenticatePactItTest {

    @Suppress("unused")
    @JvmField
    @TestTarget
    final val target: Target = HttpTarget(port = 8601, insecure = true)

    @Autowired
    private lateinit var userRepository: UserRepository

    @Before
    fun setUp() {
        userRepository.deleteAll()
    }

    @State("user exists")
    fun `user exists`() {
        userRepository.save(QuokiUser(userName = "alice", password = "s3cr3t"))
    }
}