# Anypay Go

Alternative implementation of the Anypay server software originally written in typescript.

This document serves as the plan for implementation, since no code is yet written

## Modules

- payments
- accounts
- prices
- webhooks
- logging
- monitoring
- plugins

## Plan Phase 1

First offer a REST API built on top of the current postgres database schema

The REST API should only authenticate users and provide access to their account data

Anything that does not actually require parsing and validating transactions

## Plan Phase 2

The core transaction parsing, validation, and confirmation logic.

This can be re-used as a standalone go package github.com/anypay/anypay-go/core

This library does not offer any sort of persistence.

## Plan Phase 3

Websockets server and webhooks for apps. Should work in-memory and with rabbitmq
