# -*- coding: utf-8 -*-
from sqlalchemy import Column, DateTime, ForeignKey
from sqlalchemy.dialects.postgresql import TEXT, BIGINT, BOOLEAN

from . import Base


class AccessToken(Base):
    "AccessToken for Artiefact Users"

    __tablename__ = 'access_token'

    token = Column(TEXT, primary_key=True, index=True)
    user_id = Column(BIGINT, nullable=False, index=True)
    generated_datetime = Column(DateTime(timezone=True), nullable=False)
    expiry_datetime = Column(DateTime(timezone=True), nullable=False)
    obtained_by = Column(TEXT, nullable=False)
    active = Column(BOOLEAN, nullable=False)


class AccessTokenUse(Base):
    "Access Token information of Token"

    __tablename__ = "token_access"

    # It would be useful to recrod information where the user used the access token to login
    # Location, device, etc. so primary key should be its own ID
    id = Column(BIGINT, primary_key=True)
    token = Column(TEXT, ForeignKey('access_token.token'), nullable=False)
    last_used_datetime = Column(DateTime(timezone=True), nullable=False)
