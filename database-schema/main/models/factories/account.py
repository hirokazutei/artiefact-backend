# -*- coding: utf-8 -*-
import random
import bcrypt

import factory

from main.models import account, session
from factory.alchemy import SQLAlchemyModelFactory


class ArtiefactUserFactory(SQLAlchemyModelFactory):
    class Meta:
        model = account.ArtiefactUser
        sqlalchemy_session = session

    password = bcrypt.hashpw("password1234".encode('utf-8'), bcrypt.gensalt(16)).decode("utf-8")
    email = factory.Faker('email')
    birthday = factory.Faker('date_time_this_decade', before_now=True, after_now=False)
    register_date = factory.Faker('date_time_this_decade', before_now=True, after_now=False)
    status = 'active'


class UserAgreementFactory(SQLAlchemyModelFactory):
    class Meta:
        model = account.UserAgreement
        sqlalchemy_session = session

    user_id = factory.SubFactory('main.models.factories.account.ArtiefactUserFactory')
    agreement_type = "General Registration"
    agreement_date = factory.Faker('date_time_this_decade', before_now=True, after_now=False)


class UsernameFactory(SQLAlchemyModelFactory):
    class Meta:
        model = account.Username
        sqlalchemy_session = session

    user = factory.SubFactory('main.models.factories.account.ArtiefactUserFactory')
    username_lower = factory.Faker('user_name', locale='en_US')
    username_raw = factory.Faker('user_name', locale='en_US')


class ProfileFactory(SQLAlchemyModelFactory):
    class Meta:
        model = account.Profile
        sqlalchemy_session = session

    user_id = factory.SubFactory('main.models.factories.acount.ArtiefactUserFactory')
    name = factory.Faker('first_name', locale='en_US')
    website = factory.Faker('user_name')
    bio = factory.Faker('sentence')
    gender = factory.LazyAttribute(lambda n: random.randint(0, 2))


def create_user():
    user = ArtiefactUserFactory.create()
    UsernameFactory.create(user=user)
    return user
