# -*- coding: utf-8 -*-
import random
import bcrypt

import factory

from main.models import account, session
from factory.alchemy import SQLAlchemyModelFactory


class UserFactory(SQLAlchemyModelFactory):
    class Meta:
        model = account.User
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

    user_id = factory.SubFactory('main.models.factories.account.UserFactory')
    agreement_type = "General Registration"
    agreement_date = factory.Faker('date_time_this_decade', before_now=True, after_now=False)


class UsernameFactory(SQLAlchemyModelFactory):
    class Meta:
        model = account.Username
        sqlalchemy_session = session

    user = factory.SubFactory('main.models.factories.account.UserFactory')
    username_lower = factory.Faker('user_name', locale='en_US')
    username_raw = factory.Faker('user_name', locale='en_US')


class ProfileFactory(SQLAlchemyModelFactory):
    class Meta:
        model = account.Profile
        sqlalchemy_session = session

    user_id = factory.SubFactory('main.models.factories.acount.UserFactory')
    name = factory.Faker('first_name', locale='en_US')
    website = factory.Faker('user_name')
    bio = factory.Faker('sentence')
    gender = factory.LazyAttribute(lambda n: random.randint(0, 2))


def create_user():
    user = UserFactory.create()
    UsernameFactory.create(user=user)
    return user
