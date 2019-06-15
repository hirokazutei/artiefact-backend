# -*- coding: utf-8 -*-
from sqlalchemy import Column, DateTime, ForeignKey
from sqlalchemy.orm import backref, relationship
from sqlalchemy.dialects.postgresql import TEXT, BIGINT

from . import Base

class RegisteredEmail(Base):
    "Email of Users"

    __tablename__ = "registered_email"

    id = Column(BIGINT, primary_key=True, index=True)
    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), index=True, unique=False, nullable=False)
    email = Column(TEXT, nullable=False)
    email_lower = Column(TEXT, nullable=False)
    last_used_datetime = Column(DateTime(timezone=True), nullable=False)
    status = Column(TEXT, nullable=False)

    user = relationship('artiefact_user', backref=backref('registered_email'))


class EmailVerificationRequest(Base):
    "Email Verification Request"

    __tablename__ = "email_verification_request"

    registered_email_id = Column(BIGINT, ForeignKey('registered_email.id'), primary_key=True)
    code = Column(TEXT, nullable=False)
    expiration_datetime = Column(DateTime(timezone=True), nullable=False)
    requested_datetime = Column(DateTime(timezone=True), nullable=False)

    registered_email = relationship('registered_email', backref=backref('email_verification_request'))


class EmailVerification(Base):
    "Email Verification"

    __tablename__ = "email_verification"

    request_id = Column(BIGINT, ForeignKey('email_verification_request.registered_email_id'), primary_key=True)
    verification_datetime = Column(DateTime(timezone=True), nullable=False)

    email_verification_request = relationship('email_verification_request', backref=backref('email_verification'))