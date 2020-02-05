package com.shyamz.provider.authenticate

sealed class AuthenticateResult {
    data class Success(val authenticated: Boolean) : AuthenticateResult()
    data class Error(val message: String) : AuthenticateResult()
}