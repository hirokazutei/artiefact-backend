# -*- coding: utf-8 -*-
from sqlalchemy import Date, Column, DateTime, ForeignKey
from sqlalchemy.orm import backref, relationship
from sqlalchemy.dialects.postgresql import TEXT, BIGINT, SMALLINT

from . import Base


class User(Base):
    """User"""

    __tablename__ = 'user'

    id = Column(BIGINT, primary_key=True)
    password = Column(TEXT, nullable=False)
    email = Column(TEXT, nullable=False)
    birthday = Column(Date, nullable=False)
    register_date = Column(DateTime(timezone=True), nullable=False)
    status = Column(TEXT, nullable=False, index=True)


class UserAgreement(Base):
    """User Agreemnt"""

    __tablename__ = 'user_agreement'

    id = Column(BIGINT, primary_key=True)
    user_id = Column(BIGINT, ForeignKey('user.id'), nullable=False)
    agreement_type = Column(TEXT, nullable=False)
    agreement_date = Column(DateTime(timezone=True), nullable=False)


class Username(Base):
    """Username"""

    __tablename__ = 'username'

    user_id = Column(BIGINT, ForeignKey('user.id'), primary_key=True)
    lower_name = Column(TEXT, nullable=False, unique=True, index=True)
    display_name = Column(TEXT, nullable=False)

    user = relationship('User', backref=backref('username', uselist=False))


class Profile(Base):
    """Profile"""

    __tablename__ = 'profile'

    user_id = Column(BIGINT, ForeignKey('user.id'), primary_key=True)
    name = Column(TEXT, nullable=True)
    website = Column(TEXT, nullable=True)
    bio = Column(TEXT, nullable=True)
    gender = Column(SMALLINT, nullable=True)

    user = relationship('User', backref=backref('profile', uselist=False))


class ProfilePicture(Base):
    """Profile Picture"""

    __tablename__ = 'profile picture'

    user_id = Column(BIGINT, ForeignKey('user.id'), primary_key=True)
    thumbnail = Column(TEXT, nullable=True)
    image = Column(TEXT, nullable=True)

    user = relationship('User', backref=backref('profile', uselist=False))
