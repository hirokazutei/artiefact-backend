# -*- coding: utf-8 -*-
from sqlalchemy import Date, Column, DateTime, ForeignKey
from sqlalchemy.orm import backref, relationship
from sqlalchemy.dialects.postgresql import TEXT, BIGINT, SMALLINT

from . import Base


class ArtiefactUser(Base):
    """Artiefact User"""

    __tablename__ = 'artiefact_user'

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
    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), nullable=False)
    agreement_type = Column(TEXT, nullable=False)
    agreement_date = Column(DateTime(timezone=True), nullable=False)


class Username(Base):
    """Username"""

    __tablename__ = 'username'

    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), primary_key=True)
    username_lower = Column(TEXT, nullable=False, unique=True, index=True)
    username_raw = Column(TEXT, nullable=False)

    user = relationship('ArtiefactUser', backref=backref('username', uselist=False))


class Profile(Base):
    """Profile"""

    __tablename__ = 'profile'

    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), primary_key=True)
    name = Column(TEXT, nullable=True)
    website = Column(TEXT, nullable=True)
    bio = Column(TEXT, nullable=True)
    gender = Column(SMALLINT, nullable=True)

    user = relationship('ArtiefactUser', backref=backref('profile', uselist=False))


class ProfilePicture(Base):
    """Profile Picture"""

    __tablename__ = 'profile_picture'

    user_id = Column(BIGINT, ForeignKey('artiefact_user.id'), primary_key=True)
    thumbnail = Column(TEXT, nullable=True)
    image = Column(TEXT, nullable=True)

    user = relationship('ArtiefactUser', backref=backref('profile_picture', uselist=False))
